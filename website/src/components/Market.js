import React from 'react';
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import {Stock} from './Stock';

import {Container, Button} from 'react-bootstrap';

export class Market extends React.Component {
    constructor(props) {
        super();
    }
    render() {
        return (
            <div>

                <br /><br />
            <Container>
                <Link to='/addHarvest'>
                    <Button>
                        Add Harvest
                    </Button>
                </Link>
                {'      '}
                <Link to='/sellFish'>
                    <Button>
                        Sell Fish
                    </Button>
                </Link>
                </Container>

                <br /><br /><br />
        
                <Stock/>
            </div>
        )
    }
}