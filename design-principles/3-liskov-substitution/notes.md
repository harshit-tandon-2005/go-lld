# Liskov Substitution Principle

This principle boils down to a simple but powerful idea: you should be able to replace a parent class with one of its subclasses without breaking your code.

In plain English, if your code works with a general “vehicle” class, it should still work just fine if you swap in a “car” or a “truck” class instead. No bugs, no weird surprises, no system crashes.

At its core, the Liskov Substitution Principle is all about trust, you should be able to replace a parent class with a child class and expect everything to still work smoothly. But to really follow LSP, there’s more than just swapping classes involved.

Refer:  `https://blog.logrocket.com/liskov-substitution-principle-lsp/`


Because Go uses implicit interface implementation, LSP violations usually show up as:
* surprising behavior (e.g., panics, partial implementations)
* methods that technically satisfy an interface but violate expectations (e.g., returning errors where none are expected, changing semantics)