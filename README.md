# annotorius-desktop

A minimalistic desktop IIIF viewer ([Openseadragon](https://github.com/openseadragon/openseadragon)) with annotation functions ([Annotorius-openseadragon](https://github.com/recogito/annotorious-openseadragon)).  
Made with [webview](https://github.com/zserge/webview), annotations are persisted locally into an SQLite database (`~/.annotorius/annotorius.db`).

Still a proof of concept, incomplete and not fully tested.

## run

```
~ go build
~ ./annotorius-desktop
```

(packaging into a single binary and multi-platforms builds will added later)