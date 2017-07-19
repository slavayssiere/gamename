import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import GoogleLogin from 'react-google-login';
 


class App extends Component {
  constructor(){
    super();

    this.state = {
      connect: false,
      gauth: null,
    }

    this.responseTrueGoogle = this.responseTrueGoogle.bind(this)
    this.responseFalseGoogle = this.responseFalseGoogle.bind(this)
  }

responseTrueGoogle(response){
  console.log(response);
  this.setState({
    connect: true,
    gauth: response,
  })
}

responseFalseGoogle(response){
  console.log(response);
  this.setState({
    connect: false,
    gauth: response,
  })
}

  render() {
    let connect_div = <div>
          <GoogleLogin
              clientId="202134378130-2p66ldaaptu31pg06g9mnegral199t6m.apps.googleusercontent.com"
              buttonText="Login"
              onSuccess={this.responseTrueGoogle}
              onFailure={this.responseFalseGoogle}
            />
        </div>
    
    if(this.state.connect === true){
      connect_div = <div>
        <p>Connect: { this.state.connect }</p>
        <p>access_token: { this.state.gauth.tokenId }</p>
      </div>
    }
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Welcome to React</h2>
        </div>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
        {connect_div}
      </div>
    );
  }
}

export default App;
