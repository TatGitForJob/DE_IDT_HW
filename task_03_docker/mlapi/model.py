from sklearn.datasets import load_iris
from sklearn.ensemble import RandomForestClassifier

def train_model():
    data = load_iris()
    X = data.data
    y = data.target

    model = RandomForestClassifier(n_estimators=100, random_state=42)
    model.fit(X, y)

    return model, data.target_names
