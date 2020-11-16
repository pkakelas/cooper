import React, { Component } from 'react';
import logo from '../cooper.png'
import SearchForm from './Search'

class Search extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div id="results">
                <table className="table" id="results-table">
                    <thead>
                        <tr>
                            <th scope="order">#</th>
                            <th scope="title">Title</th>
                            <th scope="url">URL</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <th scope="row">1</th>
                            <td>Github</td>
                            <td>https://github.com</td>
                        </tr>
                        <tr>
                            <th scope="row">2</th>
                            <td>Google</td>
                            <td>https://google.com</td>
                        </tr>
                        <tr>
                            <th scope="row">3</th>
                            <td>Twitter</td>
                            <td>https://twitter.com</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        );
    }
}

export default Search