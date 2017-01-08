import Inferno from 'inferno';
import Component from 'inferno-component';

export class FormComponent extends Component{

 constructor() {
    super();
    this.state = {
      username: null,
      errors: null
    };
    this.updateUsername = this.updateUsername.bind(this);
    this.submit = this.submit.bind(this);
  }

  updateUsername(event) {
    this.setState({username: event.target.value});
  }

  submit(event) {
    event.preventDefault();
    // handle form submission
    const req = new Request("/api/check", { method: "POST", body: `{"username": "${this.state.username}", "timestamp": "${Date.now().toString()}"}` });
    fetch(req)
      .then(res => res.json())
      .then(res => {
        let result = res.password ?
          `Your new password is: ${res.password}` :
          "Sorry! You don't get a password";
        this.setState({ result });
      })
      .catch(err => {
        let errors = err instanceof SyntaxError ?
          "Server found a broken" :
          err.toString();
        this.setState({ errors });
      });
  }

  render() {
    return (
      <div>
        <h2>{this.state.result}</h2>
        <form action="#" onSubmit={this.submit}>
          <div className="field">
            <label htmlFor="username">Enter Your Username</label>
            <input id="username" type="text" name="username" onChange={this.updateUsername} />
          </div>
          <div className="field">
            <button type="submit">Get Password</button>
          </div>
        </form>
        <div className="error">{this.state.errors}</div>
      </div>
    );
  }
}
