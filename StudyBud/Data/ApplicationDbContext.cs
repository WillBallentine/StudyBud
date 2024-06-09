using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore;
using StudyBud.Models;

namespace StudyBud.Data;

public class ApplicationDbContext : IdentityDbContext
{

    public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options)
        : base(options)
    {
    }

    public DbSet<User> Users { get; set; }

    public DbSet<Assignment> Assignments { get; set; }

    public DbSet<Book> Books { get; set; }

    public DbSet<Cohort> Cohorts { get; set; }

    public DbSet<Degree> Degrees { get; set; }

    public DbSet<Professor> Professors { get; set; }

    public DbSet<School> Schools { get; set; }

    public DbSet<Syllabus> Syllabi { get; set; }

}

