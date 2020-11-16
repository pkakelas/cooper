import logo from '../cooper.png'
import Search from './Search'
import Results from './Results'
import React, { Component } from 'react';

class Main extends Component {
    constructor(props) {
      super(props);

      const searchTerm = this.getSearchQueryParam();
      this.state = searchTerm && searchTerm.length != 0 ?
        {results: true, query: searchTerm} :
        {results: false, query: undefined};
    
      this.handleChange = this.handleChange.bind(this);
      this.handleSubmit = this.handleSubmit.bind(this);
    }

    getSearchQueryParam() {
      return new URLSearchParams(window.location.search).get("search");
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
        let mainComponent = this.state.results ?
            <Results query={this.state.query}></Results> :
            <Search></Search>
        let title = this.state.results ?
            `Search results for query: "${this.state.results}"` :
            "Welcome to Cooper search"
            
        return (
            <div class="App">
                <img src={logo} className="App-logo" alt="logo" />
                <h1>{title}</h1>
                {mainComponent}
            </div>
        )
    }
  }

  export default Main