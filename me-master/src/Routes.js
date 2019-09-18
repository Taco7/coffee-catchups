import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import Past from './components/past';  // or whatever the location is

export default () => (
<BrowserRouter>
    <Switch>
      <Route path="/past" component={Past}/>
    </Switch>
</BrowserRouter>
);