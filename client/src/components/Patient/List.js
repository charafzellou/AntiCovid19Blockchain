import React, {useState, useEffect} from "react";

export default function List() {
  const [patients, setPatients] = useState([]);
  const [open, setOpen] = useState();

  useEffect(() => {
  fetch('http://localhost:3000/patients')
    .then(response => response.json())
    .then(data => setPatients(data));
  }, [] );
  
  return <ul>
    {patients.map(patient =>
	<li onClick={() => setOpen(patient)}>
	{patient.patient_name}
	{open.patient_id === <span>{patient.patient_id}</span>}
	</li>)}
    </ul>;
}
