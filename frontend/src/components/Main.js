import logo from '../cooper.png'
import SearchForm from './SearchForm'
import React, { Component } from 'react';

class Main extends Component {
    constructor(props) {
      super(props);
      this.state = {value: ''};
  
      this.handleChange = this.handleChange.bind(this);
      this.handleSubmit = this.handleSubmit.bind(this);
    }
  
    handleChange(event) {
      this.setState({value: event.target.value});
    }
  
    handleSubmit(event) {
      //alert('A name was submitted: ' + this.state.value);
      window.location.href = `/?search=${this.state.value}`;
      event.preventDefault();
    }

    render() {
        return (
            <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" />
                <p> Welcome to the Cooper search! </p>
                <SearchForm></SearchForm>
            </header>
            </div>
        );
    }
  }

  export default Main