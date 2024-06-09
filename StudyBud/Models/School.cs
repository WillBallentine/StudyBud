namespace StudyBud.Models;

public class School
{
    public string SchoolId { get; set; }

    public string Name { get; set; }

    public string Address { get; set; }

    public bool Online { get; set; }

    public bool InPerson { get; set; }

    public bool Hybrid { get; set; }

    public bool CurrentlyEnrolled { get; set; }

    public int Year { get; set; }

    public string DegreeTypeInProgress { get; set; }

}

