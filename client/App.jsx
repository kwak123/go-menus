import React from 'react';
import ReactDom from 'react-dom';
import Main from './components/Main.jsx';

class App extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <Main />
    );
  }
}

export default App;
