import React from "react";

import { Navbar } from "react-bootstrap";

export class Header extends React.Component {
  constructor(props) {
    super();
  }

  render() {
    return (
      <div>
        <Navbar bg="dark" variant="dark">
          <Navbar.Brand href="/">
            <img
              alt=""
              src="http://pngimg.com/uploads/fish/fish_PNG25131.png"
              width="30"
              height="30"
              className="d-inline-block align-top"
            />
            {"The Fish Company"}
          </Navbar.Brand>
        </Navbar>
      </div>
    );
  }
}
