# tviewplus

This package extends the Terminal application framework https://github.com/rivo/tview.

## types

### FocusGroup

A FocusGroup can be used to group multiple editable views with respect to getting focus. You can cycle through members of this group using Tab, Enter and BackTab.

### StringHolder

A StringHolder adds simple notification to dependents of interface `StringChangeHandler`. 
When setting the value (`string`) all dependents (functions) are called. Typically a StringHolder can be used a Model to decouple the View from the Application (MVP pattern).

### StringListSelectionHolder

A StringListSelectionHolder adds simple notification to dependents of interface `SelectionChangeHandler`.
Typically a StringListSelectionHolder is used as the model for a DropDown.

### BoolHolder 
A BoolHolder adds simple notification to dependents of interface `BoolChangeHandler`. 
Typically a BoolHolder is used as the model for a Checkbox.

See test folder a sample program using all extensions.

## test screenshot

![test.png](test.png)

&copy; 2020 <a href="http://ernestmicklei.com">ernestmicklei.com</a>