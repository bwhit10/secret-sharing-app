from flask import Flask, render_template, request

app = Flask(__name__)


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/save', methods=['POST'])
def save():
    # TODO: save request.form to display later
    return "TODO"

if __name__ == '__main__':
    app.run()
