export async function getHealth() {
  const res = await fetch("/api/health");
  return res.json();
}

export async function getChemicals() {
  const res = await fetch("/api/chemicals");
  return res.json();
}
// src/api.js

export async function getChemical(name) {
  const res = await fetch(`/api/chemicals?name=${encodeURIComponent(name)}`);
  return await res.json();
}
