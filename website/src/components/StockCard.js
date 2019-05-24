import React from "react";

import { Card, ListGroup, ListGroupItem } from "react-bootstrap";

export const StockCard = props => {
  return (
    <ul>
      {props.stock &&
        props.stock.map((fish, index) => {
          return (
            <li key={index}>
              <Card
                style={{
                  width: "18rem"
                }}
              >
                <Card.Img variant="top" src={fish.image} />
                <Card.Body>
                  <Card.Title>{fish.species}</Card.Title>
                </Card.Body>
                <ListGroup className="list-group-flush">
                  <ListGroupItem>Amount: {fish.amount} </ListGroupItem>
                  <ListGroupItem>
                    Price: {fish.price} Algo/{fish.species}{" "}
                  </ListGroupItem>
                </ListGroup>
              </Card>
              <br />
            </li>
          );
        })}
    </ul>
  );
};
