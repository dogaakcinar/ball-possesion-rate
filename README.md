# Football Possession Calculator

This is a simple Go program that calculates the possession rates of two football teams during a match. The program allows you to switch possession between the teams and provides real-time updates of the possession rates. Additionally, an API is provided to access the possession rates programmatically.

## How to Use

1. Start the game by pressing Enter.
2. Enter the name of Team 1 and press Enter.
3. Enter the name of Team 2 and press Enter.
4. During the match, press Enter to switch possession between the teams.
5. Press 'q' followed by Enter to quit the game.

## Example Output
Possession: Team 1 - 50.00% | Team 2 - 50.00%
Team with possession: Team 1
Press Enter to switch possession or 'q' to quit:


## Workflow

The code includes a GitHub Actions workflow file (`.github/workflows/build.yaml`) that builds the code whenever changes are pushed to the repository.

## How to Run Locally

To run the program locally, follow these steps:

1. Clone the repository to your local machine.
2. Make sure you have Go installed on your machine.
3. Install the Gin dependency by running the following command in your terminal:

   ```shell
   go get github.com/gin-gonic/gin
4.Navigate to the project directory.
5.Run the following command to start the program:
go run main.go

### API Usage
You can retrieve the possession rates programmatically using the provided API. Here's an example using cURL: 
curl http://localhost:8080/api/possession

bash
Copy code
curl http://localhost:8080/api/possession

## License
This project is licensed under the MIT License. See the [LICENSE.txt](LICENSE) file for details.