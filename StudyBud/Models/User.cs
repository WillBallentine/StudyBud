using Microsoft.AspNetCore.Identity;

namespace StudyBud.Models;

public class User : IdentityUser
{
    public string ?Name { get; set; }

    public string ?FName { get; set; }

    public string ?MInitial { get; set; }

    public string ?LName { get; set; }

    public string ?Address { get; set; }

    public float ?GPA { get; set; }

    public bool ?Subscribed { get; set; }

    public List<Degree> ?Degrees { get; set; }

    public List<School> ?Schools { get; set; }

    public List<Cohort> ?Cohorts { get; set; }

    public List<Syllabus> ?Syllabi {get; set;}
}

