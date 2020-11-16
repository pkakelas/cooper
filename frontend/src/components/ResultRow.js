import React, { Component } from 'react';

class ResultRow extends Component {
    render() {
        return (
            <tr>
                <th scope="row">{this.props.id}</th>
                <td>{this.props.title}</td>
                <td>{this.props.url}</td>
            </tr>
        );
    }
}

export default ResultRow