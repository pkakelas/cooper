import React, { Component } from 'react';
import ResultRow from './ResultRow'

class Search extends Component {
    constructor(props) {
        super(props);

        this.state = {}
        this.state.a = [
            {id: "1", title: "your title", url: "https://github.com"},
            {id: "1", title: "your title", url: "https://github.com"},
            {id: "1", title: "your title", url: "https://github.com"},
        ]
    }

    async componentDidMount() {
        /*
        const response = await axios.get("http://localhost:3000", {
            params: { query: this.props.query }
        })
        */
    }

    createRows(results) {
        return results.map((r) => {
            return <ResultRow id={r.id} title={r.title} url={r.url}></ResultRow>
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
                        {this.createRows(this.state.a)}
                    </tbody>
                </table>
            </div>
        );
    }
}

export default Search