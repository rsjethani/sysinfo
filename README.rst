sysinfo
=========
An extensible Go library for getting information about installed hardware and software

Basic Usage Example
-------------------
::

  // info is a list of all battery objects found
  info, err := sysinfo.GetInfo("hardware", "battery")
  
  // Get information about the first battery object
  p, err := info.Attribute(0, "CAPACITY") // returns interface{}, error
  perc, err := info.(uint)  // extract vdesired value from the interface ref
  
  // Use it
  if perc < 20 {
      fmt.println("Battery Critical")
  }
  
**More docs coming soon**
