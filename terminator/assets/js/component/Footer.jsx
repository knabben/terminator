import React from 'react'
import FilterLink from '../containers/FilterLink'

// Manage actions on redux to edit params
// or delete the item
const Footer = () => (
  <p>
    <FilterLink filter="SHOW_ALL">
      Show All
    </FilterLink>
    {" , "}
    <FilterLink filter="SHOW_NAMESPACE">
      Show Namespace
    </FilterLink>
  </p>
)

export default Footer
