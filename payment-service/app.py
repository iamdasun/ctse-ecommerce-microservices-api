from flask import Flask, jsonify, request
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///payments.db"
app.config["SQLALCHEMY_TRACK_MODIFICATIONS"] = False

db = SQLAlchemy(app)


# Define Payment model
class Payment(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    amount = db.Column(db.Float, nullable=False)
    status = db.Column(db.String(50), nullable=False)

    def to_dict(self):
        return {"id": self.id, "amount": self.amount, "status": self.status}


# Initialize database and seed data
def initialize_database():
    with app.app_context():
        db.create_all()

        count = Payment.query.count()
        if count == 0:
            db.session.add(Payment(amount=100.00, status="SUCCESS"))
            db.session.add(Payment(amount=250.00, status="SUCCESS"))
            db.session.add(Payment(amount=75.50, status="SUCCESS"))
            db.session.commit()


@app.route("/payments", methods=["GET"])
def get_payments():
    payments = Payment.query.all()
    return jsonify([payment.to_dict() for payment in payments]), 200


@app.route("/payments/<int:payment_id>", methods=["GET"])
def get_payment(payment_id):
    payment = Payment.query.get(payment_id)

    if not payment:
        return jsonify({"error": "Payment not found"}), 404

    return jsonify(payment.to_dict()), 200


@app.route("/payments/process", methods=["POST"])
def process_payment():
    data = request.get_json()

    if not data or "amount" not in data:
        return jsonify({"error": "Amount is required"}), 400

    payment = Payment(amount=data["amount"], status="SUCCESS")
    db.session.add(payment)
    db.session.commit()

    return jsonify(payment.to_dict()), 201


if __name__ == "__main__":
    initialize_database()
    app.run(host="0.0.0.0", port=8083)
