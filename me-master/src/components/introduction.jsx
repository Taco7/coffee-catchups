import React, { Component } from 'react'

export default class Introduction extends Component {

  constructor() {
    super()
    this.state = {showImage: true}
  }

  schedule = () => {
    this.setState({showImage: !this.state.addClass});
  
  };

  render() {
    return (
      <div>
        <section id="colorlib-hero" className="js-fullheight" data-section="home">
        <div className="colorlib-narrow-content">
        <a href="" onClick={(e) => this.schedule(e)} className="blog-img"><img src="images/random.png" className="img-responsive" /><span>{this.state.title}</span></a>
        </div>
        </section>
      </div>
    )
  }
}
