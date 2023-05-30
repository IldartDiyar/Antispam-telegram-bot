import pickle
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.ensemble import RandomForestClassifier

from flask import Flask, request, jsonify

# Load the vectorizer and model
model = pickle.load(open('random_forest_model.pkl', 'rb'))
vectorizer = pickle.load(open('vectorizer.pkl', 'rb'))
app = Flask(__name__)

@app.route("/evaluate_text", methods=["POST"])
def process_text():
    data = request.get_json()
    text = data.get("text")
    print(text)
    input_vector = vectorizer.transform([text])
    pred = model.predict(input_vector)
    print(pred)

    return jsonify(pred.tolist())

if __name__ == "__main__":
    app.run()
