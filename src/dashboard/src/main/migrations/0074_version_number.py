"""Migrate the Archivematica agent version string to Archivematica-1.11."""

from django.db import migrations


def data_migration_down(apps, _):
    """Undo the migration if requested."""
    agent = apps.get_model("main", "Agent")
    agent.objects.filter(
        identifiertype="preservation system", name="Archivematica"
    ).update(identifiervalue="Archivematica-1.10")


def data_migration_up(apps, _):
    """Update the application agent version in the database."""
    agent = apps.get_model("main", "Agent")
    agent.objects.filter(
        identifiertype="preservation system", name="Archivematica"
    ).update(identifiervalue="Archivematica-1.11")


class Migration(migrations.Migration):
    """Run the migration to update the Archivematica agent version string."""

    dependencies = [("main", "0073_index_files_runsql")]

    operations = [migrations.RunPython(data_migration_up, data_migration_down)]
