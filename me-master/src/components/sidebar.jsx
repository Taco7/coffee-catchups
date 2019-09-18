import React, { Component } from 'react'
import Routes from '../Routes';


export default class Sidebar extends Component {
  render() {
    return (
      <div>
        <div>
          <nav href="#navbar" className="js-colorlib-nav-toggle colorlib-nav-toggle" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar"><i /></nav>
          <aside id="colorlib-aside" className="border js-fullheight">
            <div className="text-center">
              <div className="author-img" style={{backgroundImage: 'url(images/coffee.png)'}} />
              <h1 id="colorlib-logo"><a href="index.html">Coffee Catchups</a></h1>
            </div>
            <nav id="colorlib-main-menu" role="navigation" className="navbar">
              <div id="navbar" className="collapse">
                <ul>
                  <li className="active"><a href="future.html">Future catchups</a></li>
                  <li><a href="previous.html">Previous catchups</a></li>
                  {/* <li Link to='/past'>Previous catchups</li> */}
                  <Routes />
                  {/*<li><a href="#" data-nav-section="projects">Projects</a></li>
                  <li><a href="#" data-nav-section="blog">Blog</a></li>*/}
                  <li><a href="mindmap.html">Connections so far</a></li>
                </ul>
              </div>
            </nav>
            <div className="colorlib-footer">
              <a href="#timeline" data-nav-section="timeline">About the idea</a>
            </div>
          </aside>
        </div>
      </div>
    )
  }
}
