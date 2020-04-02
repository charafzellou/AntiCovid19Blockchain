import React, {useState, useEffect} from "react";

export default function List() {
  const [treatments, setTreatments] = useState([]);
  const [open, setOpen] = useState();

  useEffect(() => {
  fetch('http://localhost:3000/treatments')
    .then(response => response.json())
    .then(data => setTreatments(data));
  }, [] );
  
  return <ul>
    {treatments.map(treatment =>
	<li onClick={() => setOpen(treatment)}>
	{treatment.treatment_activecomponent}
	{open.treatment_id === <span>{treatment.treatment_id}</span>}
	</li>)}
    </ul>;
}
