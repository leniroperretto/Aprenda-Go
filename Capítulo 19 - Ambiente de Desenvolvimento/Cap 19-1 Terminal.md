## - Terminologia:
    - GUI: Graphical User Interface
    - CLI: Command Line Interface
        - Terminal, console, etc
    - Unix, Linux, Mac:
        - Shell, bash
    - Windows:
        Command prompt, cmd, dos prompt, powershell
## - Shell/bash commands:
    - pwd
    - ls
        - ls -la
        - Permissions: owner, group, world
        - r, w, x → 4, 2, 1 (d = directory)
        - rwxrwxrwx = owner, group, world
    - touch
    - clear
    - chmod
        - chmod options permissions filename
        - chmod 777 arquivo.ext
    - cd 
        - cd ../
        - cd qualquer/pasta/
    - env
    - rm [arquivo]
        - rm -rf [arquivo]
    - .bash_profile & .bashrc
        - .bash_profile is executed for login shells, while .bashrc is executed for interactive non-login shells. 
        - When you login (type username and password) via console, either sitting at the machine, or remotely via ssh: .bash_profile is executed to configure your shell before the initial command prompt.
    - nano [arquivo]
    - cat [arquivo]
    - grep
        - cat temp2.txt | grep enter

        - ls | grep -i documents