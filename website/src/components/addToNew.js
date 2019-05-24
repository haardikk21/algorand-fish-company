import React from "react";

import { Form, Button, Col } from "react-bootstrap";

export class addToNew extends React.Component {
  constructor(props) {
    super();
  }

  submitAddStock(event) {
    event.preventDefault();
    const data = new FormData(event.target);
    var body = {
      species: data.get("name"),
      amount: parseInt(data.get("amount")),
      price: parseInt(data.get("price")),
      image: data.get("image")
    };

    console.log(body);

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
        <Form onSubmit={e => this.submitAddStock(e)}>
          <Form.Row>
            <Form.Group as={Col} controlId="formGridName">
              <Form.Label>Name</Form.Label>
              <Form.Control
                name="name"
                id="name"
                placeholder="Enter the fish name"
              />
            </Form.Group>

            <Form.Group as={Col} controlId="formGridImage">
              <Form.Label>Image</Form.Label>
              <Form.Control
                name="image"
                id="image"
                placeholder="Enter the fish image URL"
              />
            </Form.Group>
          </Form.Row>
          <Form.Row>
            <Form.Group as={Col} controlId="formGridAmount">
              <Form.Label>Amount</Form.Label>
              <Form.Control
                name="amount"
                id="amount"
                placeholder="Enter the amount of fishes to add"
              />
            </Form.Group>

            <Form.Group as={Col} controlId="formGridPrice">
              <Form.Label>Price</Form.Label>
              <Form.Control
                name="price"
                id="price"
                placeholder="Enter the price of this fish"
              />
            </Form.Group>
          </Form.Row>
          <Button variant="primary" type="submit">
            Submit
          </Button>
        </Form>
      </div>
    );
  }
}
