Gegegeben ist:
[
  { time: 1709474171278, type: "joined", count 3 },
  { time: 1709474189660, type: "left", count 2 },
  { time: 1709474108675, type: "joined", count 4 },
  // ...
]

Gesucht ist:
  Eine Funktion, die aus einer Liste vno Events die maximale Auslastung ermittelt,
  und die Anzahl der Fahrgäste sowei den zugehörigen Zeitraum zurückgibt.

Beispiel:
  getMaxumuLoad(events)
// => { count: 8, from: ...m until: ... }

