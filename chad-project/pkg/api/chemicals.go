package api

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strings"

    "github.com/yourusername/chad-project/pkg/db"
)

// Chemical info structure
type Chemical struct {
    Name   string  `json:"name"`
    CAS    string  `json:"cas"`
    MW     float64 `json:"mw"`
    BP     float64 `json:"bp"`      // Boiling point
    Hazard string  `json:"hazard"`
}

// Helper: fetch from PubChem by name
func fetchChemicalFromPubChem(name string) (*Chemical, error) {
    Endpoint := "https://pubchem.ncbi.nlm.nih.gov/rest/pug/compound/name/%s/JSON"
    url := fmt.Sprintf(Endpoint, url.PathEscape(name))
    resp, err := http.Get(url)
    if err != nil || resp.StatusCode != 200 {
        return nil, fmt.Errorf("could not fetch from PubChem")
    }
    defer resp.Body.Close()

    // Simplified extraction (for demo: proper error/checking needed)
    var result struct {
        PC_Compounds []struct {
            Props []struct {
                URN struct {
                    Label string `json:"label"`
                } `json:"urn"`
                Value struct {
                    Sval string  `json:"sval,omitempty"`
                    Fval float64 `json:"fval,omitempty"`
                } `json:"value"`
            } `json:"props"`
        } `json:"PC_Compounds"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    chem := &Chemical{Name: name}
    for _, prop := range result.PC_Compounds[0].Props {
        key := strings.ToLower(prop.URN.Label)
        switch key {
        case "molecular weight":
            chem.MW = prop.Value.Fval
        case "cas":
            chem.CAS = prop.Value.Sval
        case "boiling point":
            chem.BP = prop.Value.Fval
        // You can expand fields here; hazard often must be synthesized from GHS/etc.
        }
    }
    chem.Hazard = "Data from PubChem" // Expand with extra API calls/GHS if necessary

    return chem, nil
}

// `/api/chemicals?name=chlorine` handler
func ChemicalHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        http.Error(w, "Missing 'name' query", http.StatusBadRequest)
        return
    }

    // 1. Try cache in SQL first
    var chem Chemical
    err := db.DB.QueryRow(
        "SELECT name, cas, mw, bp, hazard FROM chemicals WHERE name=@p1", name,
    ).Scan(&chem.Name, &chem.CAS, &chem.MW, &chem.BP, &chem.Hazard)

    if err == nil {
        json.NewEncoder(w).Encode(chem)
        return
    } else if err != sql.ErrNoRows {
        http.Error(w, "DB error", 500)
        return
    }

    // 2. If not in DB, fetch from PubChem
    pubChem, err := fetchChemicalFromPubChem(name)
    if err != nil {
        http.Error(w, "Could not fetch from PubChem", http.StatusServiceUnavailable)
        return
    }
    // 3. Insert into DB for next time
    _, _ = db.DB.Exec(
        "INSERT INTO chemicals (name, cas, mw, bp, hazard) VALUES (@p1, @
