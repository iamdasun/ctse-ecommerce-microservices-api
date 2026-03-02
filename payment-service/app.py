from flask import Flask, jsonify, request

app = Flask(__name__)

# In-memory data store
payments = []
next_id = 1

@app.route('/payments', methods=['GET'])
def get_payments():
    return jsonify(payments), 200

@app.route('/payments/<int:payment_id>', methods=['GET'])
def get_payment(payment_id):
    for payment in payments:
        if payment['id'] == payment_id:
            return jsonify(payment), 200
    return jsonify({'error': 'Payment not found'}), 404

@app.route('/payments/process', methods=['POST'])
def process_payment():
    global next_id

    data = request.get_json()
    if not data or 'amount' not in data:
        return jsonify({'error': 'Amount is required'}), 400

    payment = {
        'id': next_id,
        'amount': data['amount'],
        'status': 'SUCCESS'
    }
    next_id += 1
    payments.append(payment)

    return jsonify(payment), 201

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8083)
