# Go Text Summarizer using Hugging Face API

This project is a Go application that uses the Hugging Face Inference API to summarize text. The application sends a POST request to the Hugging Face API with the input text and prints the summarized text returned by the model.

## Tech Stack

- Go
- Hugging Face Inference API

## Setup

1. **Clone the repository**:

   ```sh
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Install Go: Follow the instructions on the official Go website to install Go on your machine.

3. Set your Hugging Face API token: Export your Hugging Face API token as an environment variable:

```
export HUGGING_FACE_API_TOKEN=your_hugging_face_api_token
```

## Usage

Run the application:

```
go run main.go -text "Your text to summarize"
```

Replace "Your text to summarize" with the actual text you want to summarize.

## Example

Here is an example command to summarize a text about Indian culture:

```
go run main.go -text "Indian culture is one of the oldest and most diverse cultures in the world. The Indian subcontinent is home to a variety of traditions, languages, religions, and social customs. The culture of India has been shaped by its long history, unique geography, and diverse demography. India is known for its rich heritage, which includes art, music, dance, festivals, and cuisine. The country has a deep-rooted tradition of spirituality and philosophy, with major religions such as Hinduism, Buddhism, Jainism, and Sikhism originating here. Indian festivals like Diwali, Holi, and Navratri are celebrated with great enthusiasm and are known for their vibrant and colorful nature. Indian cuisine is famous for its use of spices and herbs, with dishes like curry, biryani, and dosa being popular worldwide. Traditional Indian art forms such as Bharatanatyam, Kathak, and Odissi dance, as well as classical music styles like Hindustani and Carnatic, are highly regarded. The Indian family structure is typically patriarchal, with a strong emphasis on family values and respect for elders. Despite modernization and globalization, many Indians continue to uphold their cultural heritage and traditions."
```

## Acknowledgements

- Hugging Face for providing the Inference API and models.
