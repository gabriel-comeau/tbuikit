## Termbox UI Kit
Termbox UI Kit is a helper library which wraps around the [Termbox-go](https://www.github.com/nsf/termbox-go) library.

It offers a few levels of abstraction to build console-based user interfaces.  The main types are the UI, which
represents the ui in its entirety.  Next is the screen, which is a container for widgets.  A UI can contain any
number of screens, of which one can be active.  Finally there are the widgets, which can be placed onto screens.
Some kinds of widgets can be "selected" and others are read only.  Additionally, some widgets are backed by
buffers, which the application can read from / write to and the contents of which will be displayed in the widget.

Widget positioning is handled by callback function (the scope of which belongs to your application).  This means you can
define either fixed positions and sizes or use resizable (by accessing the console's width and height).

Keybindings are handled similarly; keys are bound to the ui, screens and widgets which take a callback function.  The functions
return an instance of whatever ui element owned the key binding, as well as the binding itself.

### IMPORTANT - WARNING
The API is still very much in flux - at the moment there are some very long constructor methods to make new widgets
and I'm still considering whether or not I'd prefer to have a bunch of setter methods instead of that.

There are also a lot of features I'd like to add, some of which could have consequences on the API.

### Installation
Install and update this go package with `go get -u github.com/gabriel-comeau/tbuikit`
You'll also want to install termbox-go (for access to the colors) with `go get -u github.com/nsf/termbox-go`

### Examples
As an example application, there's a simple chat client (and server to go with it) which I've written
at [SimpleChatClient](https://www.github.com/gabriel-comeau/SimpleChatClient) and [SimpleChatServer](https://www.github.com/gabriel-comeau/SimpleChatServer)
