import React, { Component } from 'react';

class Search extends Component {
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
        <form className="container-sm" id="search-bar" onSubmit={this.handleSubmit}>
          <div className="form-group">
            <input type="text" className="form-control" onChange={this.handleChange} placeholder="Doggies like cooper.. 🐕 "/>
          </div>
          <input type="submit" className="btn btn-default" value="Search" />
        </form>
      );
    }
  }

  export default Search