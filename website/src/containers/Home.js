import React from 'react';
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

import { Market } from './../components/Market';
import { addHarvest } from './../components/addHarvest';
import { sellFish } from './../components/sellFish';
import { Header } from './../components/Header';
import {addToExisting} from '../components/addToExisting';
import { addToNew } from '../components/addToNew';

export default class Home extends React.Component{
    constructor(props){
        super();
    }
    
    render(){
        return(
            <Router>
                <div>
                    <Header/>
                    <Route path="/" exact="exact" component={Market} />
                    <Route path="/addHarvest" component={addHarvest} />
                    <Route path="/sellFish" component={sellFish} />
                    <Route path="/addToExisting" component={addToExisting} />
                    <Route path="/addToNew" component={addToNew} />
                </div>
            </Router>
        )
    }
}