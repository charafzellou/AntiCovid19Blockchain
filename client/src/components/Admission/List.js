import React, {useState, useEffect} from "react";

export default function List() {
  const [admissions, setAdmissions] = useState([]);
  const [open, setOpen] = useState();

  useEffect(() => {
  fetch('http://localhost:3000/admissions')
    .then(response => response.json())
    .then(data => setAdmissions(data));
  }, [] );
  
  return <ul>
    {admissions.map(admission =>
	<li onClick={() => setOpen(admission)}>
	{admission.Title}
	{open.patient_id && open.treatment_id === admission.patient_id && <span>{admission.treatment_id}</span>}
	</li>)}
    </ul>;
}
