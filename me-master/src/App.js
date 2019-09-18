import React, { Component } from 'react';
import './App.css';
import Past from './components/past'
import Future from './components/future'
import Introduction from './components/introduction'
import Sidebar from './components/sidebar'
import Timeline from './components/timeline'
import ReactDOM from 'react-dom'
import Routes from './Routes';

class App extends Component {
  render() {
    return (
      <div id="colorlib-page">
        <div id="container-wrap">
         	<Sidebar></Sidebar>
				<div id="colorlib-main">
					<Introduction></Introduction>
					{/* <About></About> */}
					<Timeline></Timeline>
          	</div>
      	</div>
      </div>
    );
  }
}



export default App;


