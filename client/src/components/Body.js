import React from "react";
import {Route, Switch} from "react-router-dom";
import Home from "./Home";
import AdmissionList from "./Admission/List";
import AdmissionDetails from "./Admission/Details";
import PatientList from "./Patient/List";
import PatientDetails from "./Patient/Details";
import TreatmentList from "./Treatment/List";
import TreatmentDetails from "./Treatment/Details";

export default function Body({trigger}) {
  return <div>
      <Switch>
        <Route path="/" component={Home} exact={true}/>
        
        <Route path="/admissions-list" component={AdmissionList}/>
        <Route path="/admissions/:title" component={AdmissionDetails}/>
        
        <Route path="/patients-list" component={PatientList}/>
        <Route path="/patients/:title" component={PatientDetails}/>
        
        <Route path="/treatments-list" component={TreatmentList}/>
        <Route path="/treatments/:title" component={TreatmentDetails}/>
      </Switch>
    <button onClick={() => trigger("Titi")}>Change Name</button>
  </div>
}
