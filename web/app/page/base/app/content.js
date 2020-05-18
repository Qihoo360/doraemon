import React, { Component } from 'react';

export default class Content extends Component {
  render() {
    return (
      <div id="main">app content
        {
          this.props.children
        }
      </div>
    )
  }
}
