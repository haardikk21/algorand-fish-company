import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

export class addHarvest extends React.Component {
  constructor(props) {
    super();
  }

  render() {
    return (
      <div>
        <center>
          <br />
          <br />
          <br />
          <br />
          <br />
          <br />
          <br />
          <br />
          <button>
            <Link
              style={{ display: "block", height: "100%" }}
              to="/addToExisting"
            >
              Add to existing fish stock
            </Link>
          </button>
          <br />
          <br />
          <br />
          <br />
          <br />
          <br />
          <button>
            <Link style={{ display: "block", height: "100%" }} to="/addToNew">
              Add new fish types
            </Link>
          </button>
        </center>
      </div>
    );
  }
}
