import React from "react";
import { Stock } from "./Stock";

import { Row, Col, Form, Button } from "react-bootstrap";

export class sellFish extends React.Component {
  constructor(props) {
    super();
    this.state = {
      code: null
    };
  }

  submitSellStock(event) {
    event.preventDefault();
    const data = new FormData(event.target);

    var addr = data.get("to");
    console.log(addr);
    var body = {
      addr: [
        {
          species: data.get("name"),
          price: parseInt(data.get("price")),
          amount: parseInt(data.get("amount"))
        }
      ]
    };

    fetch("http://localhost:5000/sell", {
      method: "POST",
      body: JSON.stringify(body),
      headers: {
        "Content-Type": "application/json"
      }
    })
      .then(res => res.json())
      .then(res => {
        this.setState({ code: res });
      });
  }

  render() {
    if (this.state.code) {
      var src = "data:image/png;base64," + this.state.code;
      return (
        <div>
          <Row>
            <Col>
              <Stock />
            </Col>
            <Col>
              <Form onSubmit={e => this.submitSellStock(e)}>
                <Form.Row>
                  <Form.Group as={Col} controlId="formGridAddress">
                    <Form.Label>To</Form.Label>
                    <Form.Control
                      name="to"
                      id="to"
                      placeholder="Sender's Algo Addres"
                    />
                  </Form.Group>
                </Form.Row>
                <Form.Row>
                  <Form.Group as={Col} controlId="formGridFish">
                    <Form.Label>Fish</Form.Label>
                    <Form.Control
                      name="name"
                      id="name"
                      placeholder="Fish name"
                    />
                  </Form.Group>
                  <Form.Group as={Col} controlId="formGridAmount">
                    <Form.Label>Amount</Form.Label>
                    <Form.Control
                      name="amount"
                      id="amount"
                      placeholder="Amount of fishes"
                    />
                  </Form.Group>
                </Form.Row>
                <Button variant="primary" type="submit">
                  Submit
                </Button>
              </Form>
              <center>
                <img id="qrCode" src={src} alt="" />
              </center>
            </Col>
          </Row>
        </div>
      );
    } else {
      return (
        <div>
          <Row>
            <Col>
              <Stock />
            </Col>
            <Col>
              <Form onSubmit={e => this.submitSellStock(e)}>
                <Form.Row>
                  <Form.Group as={Col} controlId="formGridAddress">
                    <Form.Label>To</Form.Label>
                    <Form.Control
                      name="to"
                      id="to"
                      placeholder="Sender's Algo Addres"
                    />
                  </Form.Group>
                </Form.Row>
                <Form.Row>
                  <Form.Group as={Col} controlId="formGridFish">
                    <Form.Label>Fish</Form.Label>
                    <Form.Control
                      name="name"
                      id="name"
                      placeholder="Fish name"
                    />
                  </Form.Group>
                  <Form.Group as={Col} controlId="formGridAmount">
                    <Form.Label>Amount</Form.Label>
                    <Form.Control
                      name="amount"
                      id="amount"
                      placeholder="Amount of fishes"
                    />
                  </Form.Group>
                </Form.Row>
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
}
