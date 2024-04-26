# Hyprland Workspace Switch

- A simple script to move between workspaces.
- Add a binding to hyprland.conf to switch workspaces with a key binding.

  Ex ($mainMod = Super Key) :-
  ```conf
  bind  = $mainMod, TAB, exec, <*ScriptPath>
  ```

> `*ScriptPath` : <br>
    Method 1 : <br>
        - Build the go script with `go build main.go` or use the already built file from releases.<br>
        - Add the path of the build file as the ScriptPath<br><br>
    Method 2 : <br>
        - Directly run the go script with `go run main.go`

> Note :- `Display` name is currently not used for anything. It will be utilized in a future iteration.