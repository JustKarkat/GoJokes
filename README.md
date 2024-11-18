# GoLang Joke Fetcher 🤣

This is a simple project written in **GoLang** to enhance my skills with the Go programming language. The program utilizes the [Official Joke API](https://official-joke-api.appspot.com) to fetch and display random jokes of various types.

## 🚀 Features

- Fetches random jokes from the [Official Joke API](https://official-joke-api.appspot.com).
- Supports different joke types.
- Lightweight and efficient GoLang implementation.

## 🛠️ Technologies Used

- **GoLang**: A statically typed, compiled language known for its simplicity and efficiency.
- **Official Joke API**: A free API for fetching jokes.

## 📦 Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/JustKarkat/GoJokes.git
    cd GoJokes
    ```

2. Build the project:
    ```bash
    go build main.go
    ```

3. Run the program:
    ```bash
    ./main
    ```

## 🌟 How It Works

The program makes an HTTP GET request to the Official Joke API endpoint to fetch the joke types. Once retrieved, it asks which joke type, and then gets a joke of that type and displays the joke in the console.

### Example Output

```plaintext
Fetching a random joke...
Q: Why don't skeletons fight each other?
A: They don't have the guts!
```
## 📖 Learning Goals

- Improve familiarity with GoLang syntax and conventions.
- Learn how to interact with external APIs using Go's `net/http` package.
- Practice structuring a Go project.

## 🛠️ API Endpoints Used

- **Random Joke**: [https://official-joke-api.appspot.com/random_joke](https://official-joke-api.appspot.com/random_joke)
- **Jokes by Type**: [https://official-joke-api.appspot.com/jokes/{type}/random](https://official-joke-api.appspot.com/jokes/{type}/random)

## 🚧 Future Enhancements

- Add a feature to get multiple jokes.
- Implement a CLI menu for easier user interaction.
- Add unit tests to ensure reliability.

## 🤝 Contributing

Contributions are welcome! Feel free to fork this repo and submit a pull request with your improvements. If you have any suggestions or bug reports, please create an issue.
(I am also trying to work on writing good ReadMe's so any help there is nice too :D!

## 📝 License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

## 📬 Contact

Created by **[Joshua Russell](https://github.com/JustKarkat)**  
[joshua@jrussell.dev](mailto:joshua@jrussell.dev)
Feel free to reach out via GitHub or email if you have questions or feedback!

---

Happy coding :D! 🎉
