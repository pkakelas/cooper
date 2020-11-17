import axios from 'axios'
import React, { Component } from 'react';
import ResultRow from './ResultRow'

class Search extends Component {
    constructor(props) {
        super(props);
        this.state = {results: []}
    }

    async componentDidMount() {
        const response = await axios.get("http://localhost:8080", {
            params: { query: this.props.query }
        })

        this.setState({results: this.createRows(response.data)})
    }

    createRows(results) {
        return results.map((r, idx) => {
            return <ResultRow id={idx + 1} title={r.title} url={r.url}></ResultRow>
        })
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
                        {this.state.results}
                    </tbody>
                </table>
            </div>
        );
    }
}

export default Search