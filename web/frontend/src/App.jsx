import React, { useState } from "react";
import { getChemical } from "./api";

export default function App() {
  const [input, setInput] = useState("");
  const [chemical, setChemical] = useState(null);

  const searchChemical = async () => {
    const data = await getChemical(input);
    setChemical(data);
  };

  return (
    <div>
      <h1>CHAD - Chemical Search</h1>
      <input
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="Enter chemical name"
      />
      <button onClick={searchChemical}>Search</button>

      {chemical && (
        <div>
          <h2>{chemical.name}</h2>
          <p>CAS: {chemical.cas}</p>
          <p>MW: {chemical.mw}</p>
          <p>Boiling Point: {chemical.bp}</p>
          <p>Hazard: {chemical.hazard}</p>
        </div>
      )}
    </div>
  );
}
