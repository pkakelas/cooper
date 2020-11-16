import React, { Component } from 'react';
import logo from '../cooper.png'
import SearchForm from './SearchForm'

class Search extends Component {
    render() {
        return (
            <div className="App">
                <header className="App-header">
                    <h1> Welcome to Cooper search</h1>
                    <img src={logo} className="App-logo" alt="logo" />
                    <SearchForm></SearchForm>
                </header>
            </div>
        );
    }
}

export default Search