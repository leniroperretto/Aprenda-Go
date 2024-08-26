## Cap. 18 – Concorrência – 3. Discussão: Condição de corrida

- Agora vamos dar um mergulho na documentação:
    - https://golang.org/doc/effective_go.h...
    - https://pt.wikipedia.org/wiki/Multipl...
    - O que é yield? runtime.Gosched()
- Race condition: 
        Função 1       var     Função 2
         Lendo: 0   →   0
         Yield          0   →   Lendo: 0
         var++: 1               Yield
         Grava: 1   →   1       var++: 1
                        1   ←   Grava: 1
         Lendo: 1   ←   1
         Yield          1   →   Lendo: 1
         var++: 2               Yield
         Grava: 2   →   2       var++: 2
                        2   ←   Grava: 2
- E é por isso que vamos ver mutex, atomic e, por fim, channels.