using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace StudyBud.Data.Migrations
{
    /// <inheritdoc />
    public partial class AllModelAdd : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "Professors",
                columns: table => new
                {
                    ProfessorId = table.Column<string>(type: "TEXT", nullable: false),
                    Name = table.Column<string>(type: "TEXT", nullable: false),
                    Email = table.Column<string>(type: "TEXT", nullable: false),
                    Phone = table.Column<string>(type: "TEXT", nullable: false),
                    School = table.Column<string>(type: "TEXT", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Professors", x => x.ProfessorId);
                });

            migrationBuilder.CreateTable(
                name: "Users",
                columns: table => new
                {
                    UserId = table.Column<string>(type: "TEXT", nullable: false),
                    Name = table.Column<string>(type: "TEXT", nullable: false),
                    FName = table.Column<string>(type: "TEXT", nullable: false),
                    LName = table.Column<string>(type: "TEXT", nullable: false),
                    MInitial = table.Column<string>(type: "TEXT", nullable: false),
                    Email = table.Column<string>(type: "TEXT", nullable: false),
                    Phone = table.Column<string>(type: "TEXT", nullable: false),
                    Address = table.Column<string>(type: "TEXT", nullable: false),
                    GPA = table.Column<float>(type: "REAL", nullable: false),
                    Subscribed = table.Column<bool>(type: "INTEGER", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Users", x => x.UserId);
                });

            migrationBuilder.CreateTable(
                name: "Cohorts",
                columns: table => new
                {
                    CohortId = table.Column<string>(type: "TEXT", nullable: false),
                    Topic = table.Column<string>(type: "TEXT", nullable: false),
                    Mentor = table.Column<bool>(type: "INTEGER", nullable: false),
                    StandardUser = table.Column<bool>(type: "INTEGER", nullable: false),
                    UserId = table.Column<string>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Cohorts", x => x.CohortId);
                    table.ForeignKey(
                        name: "FK_Cohorts_Users_UserId",
                        column: x => x.UserId,
                        principalTable: "Users",
                        principalColumn: "UserId");
                });

            migrationBuilder.CreateTable(
                name: "Schools",
                columns: table => new
                {
                    SchoolId = table.Column<string>(type: "TEXT", nullable: false),
                    Name = table.Column<string>(type: "TEXT", nullable: false),
                    Address = table.Column<string>(type: "TEXT", nullable: false),
                    Online = table.Column<bool>(type: "INTEGER", nullable: false),
                    InPerson = table.Column<bool>(type: "INTEGER", nullable: false),
                    Hybrid = table.Column<bool>(type: "INTEGER", nullable: false),
                    CurrentlyEnrolled = table.Column<bool>(type: "INTEGER", nullable: false),
                    Year = table.Column<int>(type: "INTEGER", nullable: false),
                    DegreeTypeInProgress = table.Column<string>(type: "TEXT", nullable: false),
                    UserId = table.Column<string>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Schools", x => x.SchoolId);
                    table.ForeignKey(
                        name: "FK_Schools_Users_UserId",
                        column: x => x.UserId,
                        principalTable: "Users",
                        principalColumn: "UserId");
                });

            migrationBuilder.CreateTable(
                name: "Syllabi",
                columns: table => new
                {
                    SyllabusID = table.Column<string>(type: "TEXT", nullable: false),
                    ClassTitle = table.Column<string>(type: "TEXT", nullable: false),
                    School = table.Column<string>(type: "TEXT", nullable: false),
                    ProfessorId = table.Column<string>(type: "TEXT", nullable: false),
                    TA = table.Column<bool>(type: "INTEGER", nullable: false),
                    CreditHours = table.Column<int>(type: "INTEGER", nullable: false),
                    StartDate = table.Column<DateTime>(type: "TEXT", nullable: false),
                    EndDate = table.Column<DateTime>(type: "TEXT", nullable: false),
                    Semester = table.Column<string>(type: "TEXT", nullable: false),
                    CourseObjectives = table.Column<string>(type: "TEXT", nullable: false),
                    Objectives = table.Column<string>(type: "TEXT", nullable: false),
                    TechRequirements = table.Column<string>(type: "TEXT", nullable: false),
                    Misc = table.Column<string>(type: "TEXT", nullable: false),
                    LateWork = table.Column<bool>(type: "INTEGER", nullable: false),
                    UserId = table.Column<string>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Syllabi", x => x.SyllabusID);
                    table.ForeignKey(
                        name: "FK_Syllabi_Professors_ProfessorId",
                        column: x => x.ProfessorId,
                        principalTable: "Professors",
                        principalColumn: "ProfessorId",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_Syllabi_Users_UserId",
                        column: x => x.UserId,
                        principalTable: "Users",
                        principalColumn: "UserId");
                });

            migrationBuilder.CreateTable(
                name: "Degrees",
                columns: table => new
                {
                    DegreeId = table.Column<string>(type: "TEXT", nullable: false),
                    DegreeType = table.Column<string>(type: "TEXT", nullable: false),
                    IssuingSchoolSchoolId = table.Column<string>(type: "TEXT", nullable: true),
                    GPA = table.Column<float>(type: "REAL", nullable: false),
                    YearStarted = table.Column<DateTime>(type: "TEXT", nullable: false),
                    YearFinished = table.Column<DateTime>(type: "TEXT", nullable: false),
                    Graduated = table.Column<bool>(type: "INTEGER", nullable: false),
                    UserId = table.Column<string>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Degrees", x => x.DegreeId);
                    table.ForeignKey(
                        name: "FK_Degrees_Schools_IssuingSchoolSchoolId",
                        column: x => x.IssuingSchoolSchoolId,
                        principalTable: "Schools",
                        principalColumn: "SchoolId");
                    table.ForeignKey(
                        name: "FK_Degrees_Users_UserId",
                        column: x => x.UserId,
                        principalTable: "Users",
                        principalColumn: "UserId");
                });

            migrationBuilder.CreateTable(
                name: "Assignments",
                columns: table => new
                {
                    AssignmentId = table.Column<string>(type: "TEXT", nullable: false),
                    Name = table.Column<string>(type: "TEXT", nullable: false),
                    Type = table.Column<string>(type: "TEXT", nullable: false),
                    Description = table.Column<string>(type: "TEXT", nullable: false),
                    PercentOfGrade = table.Column<decimal>(type: "TEXT", nullable: false),
                    SyllabusID = table.Column<string>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Assignments", x => x.AssignmentId);
                    table.ForeignKey(
                        name: "FK_Assignments_Syllabi_SyllabusID",
                        column: x => x.SyllabusID,
                        principalTable: "Syllabi",
                        principalColumn: "SyllabusID");
                });

            migrationBuilder.CreateTable(
                name: "Books",
                columns: table => new
                {
                    BookId = table.Column<string>(type: "TEXT", nullable: false),
                    Title = table.Column<string>(type: "TEXT", nullable: false),
                    ISBN = table.Column<string>(type: "TEXT", nullable: false),
                    Author = table.Column<string>(type: "TEXT", nullable: false),
                    Description = table.Column<string>(type: "TEXT", nullable: false),
                    Length = table.Column<int>(type: "INTEGER", nullable: false),
                    SyllabusID = table.Column<string>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Books", x => x.BookId);
                    table.ForeignKey(
                        name: "FK_Books_Syllabi_SyllabusID",
                        column: x => x.SyllabusID,
                        principalTable: "Syllabi",
                        principalColumn: "SyllabusID");
                });

            migrationBuilder.CreateIndex(
                name: "IX_Assignments_SyllabusID",
                table: "Assignments",
                column: "SyllabusID");

            migrationBuilder.CreateIndex(
                name: "IX_Books_SyllabusID",
                table: "Books",
                column: "SyllabusID");

            migrationBuilder.CreateIndex(
                name: "IX_Cohorts_UserId",
                table: "Cohorts",
                column: "UserId");

            migrationBuilder.CreateIndex(
                name: "IX_Degrees_IssuingSchoolSchoolId",
                table: "Degrees",
                column: "IssuingSchoolSchoolId");

            migrationBuilder.CreateIndex(
                name: "IX_Degrees_UserId",
                table: "Degrees",
                column: "UserId");

            migrationBuilder.CreateIndex(
                name: "IX_Schools_UserId",
                table: "Schools",
                column: "UserId");

            migrationBuilder.CreateIndex(
                name: "IX_Syllabi_ProfessorId",
                table: "Syllabi",
                column: "ProfessorId");

            migrationBuilder.CreateIndex(
                name: "IX_Syllabi_UserId",
                table: "Syllabi",
                column: "UserId");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "Assignments");

            migrationBuilder.DropTable(
                name: "Books");

            migrationBuilder.DropTable(
                name: "Cohorts");

            migrationBuilder.DropTable(
                name: "Degrees");

            migrationBuilder.DropTable(
                name: "Syllabi");

            migrationBuilder.DropTable(
                name: "Schools");

            migrationBuilder.DropTable(
                name: "Professors");

            migrationBuilder.DropTable(
                name: "Users");
        }
    }
}
