import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import Past from './components/past'
import Future from './components/future'

import Introduction from './components/introduction'
import { Route, Link, BrowserRouter as Router } from 'react-router-dom'


const routing = (
    <Router>
   <div>
        <Route exact path="/" component={Introduction} />
        <Route path="/past" component={Past} />
        <Route path="/future" component={Future} />
      </div>
    </Router>
  )

ReactDOM.render(<App />, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();
