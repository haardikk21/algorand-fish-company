import React from "react";
import { Stock } from "./Stock";

import { Row, Col, Form, Button } from "react-bootstrap";

export class addToExisting extends React.Component {
  constructor(props) {
    super();
  }

  submitAddStock(event) {
    event.preventDefault();
    const data = new FormData(event.target);
    var body = {
      species: data.get("name"),
      amount: parseInt(data.get("amount"))
    };

    fetch("http://localhost:5000/add", {
      method: "POST",
      body: JSON.stringify(body),
      headers: {
        "Content-Type": "application/json"
      },
      mode: "no-cors"
    });
  }

  render() {
    return (
      <div>
        <Row>
          <Col>
            <Stock />
          </Col>
          <Col>
            <Form onSubmit={e => this.submitAddStock(e)}>
              <Form.Group as={Col} controlId="formGridFish">
                <Form.Label>Fish</Form.Label>
                <Form.Control name="name" id="name" placeholder="Fish name" />
              </Form.Group>

              <Form.Group as={Col} controlId="formGridAmount">
                <Form.Label>Amount</Form.Label>
                <Form.Control
                  name="amount"
                  id="amount"
                  placeholder="Amount of fishes"
                />
              </Form.Group>

              <Button variant="primary" type="submit">
                Submit
              </Button>
            </Form>
          </Col>
        </Row>
      </div>
    );
  }
}
