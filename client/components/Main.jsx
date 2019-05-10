import React from 'react';
import T from 'prop-types';

class Main extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <h1> Welcome! </h1>
      </div>
    );
  }
}

Main.propTypes = {
  // Fill in your types here
};

export default Main;