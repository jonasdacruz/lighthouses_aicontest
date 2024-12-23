# Go Bot Flow
- Initialize Bot configuration:
    - `--bn`: Bot name
    - `--la`: Bot listen address
    - `--gs`: Game server address
- Connect to the Game server and Join the current game
    - Sends the `BotName` and the `BotAddress` to the Game server
    - Receives the `PlayerID` from the Game server
- Start the two main services to play the game
    - `InitialState`: Receives the initial state from the Game server with the following information
        - `PlayerID`
        - `PlayerCount`
        - `Position`
        - `Map`
        - `Lighthouses`
    - `Turn`: Sends the turn action to the Game server
        - Receives from the Game server the following information
            - `Position`
            - `Score`
            - `Energy`
            - `View`
            - `Lighthouses`
        - Returns the action to be executed on the current turn to the Game server
            - `Action`
            - `Destination`
            - `Energy`
