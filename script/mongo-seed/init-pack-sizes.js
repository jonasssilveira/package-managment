db = db.getSiblingDB('packs_db');

db.pack_sizes.insertMany([
    { size: 250 },
    { size: 500 },
    { size: 1000 },
    { size: 2000 },
    { size: 5000 }
]);