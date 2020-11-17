import React, { Component } from 'react';

class ResultRow extends Component {
    render() {
        return (
            <tr>
                <th scope="row">{this.props.id}</th>
                <td>{this.props.title}</td>
                <td><a href={this.props.url}>{this.props.url}</a></td>
            </tr>
        );
    }
}

export default ResultRow