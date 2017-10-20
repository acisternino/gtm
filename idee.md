# Idee per una migliore add

Il Frame guadagna un pointer al parent. E' piu' grande ma ne guadagniamo
in semplicita'.

```go
type Frame struct {
    Address uint32
    Length  uint32
    parent  *Frame
    left    *Frame
    right   *Frame
}
```

La `Add()` diventa ricorsiva, non dovrebbero esserci problemi anche se la
rebalance diventa critica perche' se chiamiamo Add() di nuovo allora e'
complesso.

```go
func (f *Frame) Add(nf *Frame) {
    if nf.Length < f.Length {
        branch := f.branchFor(new)
        branch.Add(new)
    } else {
        // arrived
        if f.parent != nil {
            f.parent.append(new) // simple append
        }
        tbr := nf.append(f)
        tbr.rebalanceOnto(nf)
    }
}
```

API:

* **Add(Frame)**: come sopra. Magari con un booleano che flagghi se ribilanciare o no.
* **append(Frame)**: prende l'argomento e lo appende direttamente a self
  sovrascrivendo uno dei due `left` o `right`. Ritorna il ramo da ribilanciare
  se esiste.
* **branchFor()**: ritorna `left` o `right` a seconda dei valori di `Address`
* **rebalanceOnto()**: prende un subtree, lo naviga depth-first post-order e
  aggiunge ogni nodo all argomento. L'ideale sarebbe che in questo caso non ci
  sia rebalance.
* **addNoBalance()**: forse fa lo stesso lavoro della Add ma senza il bilanciamento.
  In altre parole naviga finche' trova un foglia e lo aggiunge li.
  Alternativo al boolean flag della Add()
