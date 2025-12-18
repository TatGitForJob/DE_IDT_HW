from fastapi import FastAPI
from pydantic import BaseModel
import psycopg2
import os
import json

from model import train_model

app = FastAPI()

model, target_names = train_model()

class PredictRequest(BaseModel):
    features: list[float]

def get_db_connection():
    return psycopg2.connect(
        host=os.getenv("DB_HOST"),
        database=os.getenv("DB_NAME"),
        user=os.getenv("DB_USER"),
        password=os.getenv("DB_PASSWORD")
    )

@app.post("/predict")
def predict(request: PredictRequest):
    prediction_idx = model.predict([request.features])[0]
    prediction = target_names[prediction_idx]

    conn = get_db_connection()
    cursor = conn.cursor()

    cursor.execute(
        "INSERT INTO predictions (features, prediction) VALUES (%s, %s)",
        (json.dumps(request.features), prediction)
    )

    conn.commit()
    cursor.close()
    conn.close()

    return {"prediction": prediction}
