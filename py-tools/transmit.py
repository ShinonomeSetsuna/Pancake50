import json

with open("./dxdata.json", "r", encoding="utf-8") as f:
    data = json.loads(f.read())
    op = {}
    songs = data["songs"]
    for song in songs:
        if song["category"] == "宴会場":
            continue
        op[song["songId"]] = {"DX": {}, "SD": {}}
        for sheet in song["sheets"]:
            op[song["songId"]][("DX", "SD")[sheet["type"] == "std"]][
                str(sheet["difficulty"]).capitalize()
            ] = sheet["internalLevelValue"]
    with open("./ds.json", "w", encoding="utf-8") as f1:
        f1.write(json.dumps(op, ensure_ascii=False))
